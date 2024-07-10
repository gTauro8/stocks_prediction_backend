import sys
import requests
import pandas as pd
import yfinance as yf
import numpy as np
from datetime import datetime, timedelta
import matplotlib.pyplot as plt
from scipy.optimize import minimize

def login(username, password):
    url = 'http://localhost:8080/login'
    response = requests.post(url, json={"username": username, "password": password})
    if response.status_code == 200:
        return response.json().get('token')
    else:
        raise Exception("Login failed")

def get_predictions(stock_tickers, start_date, end_date, days_in_future):
    url = 'http://localhost:5000/predict'
    data = {
        "stock_tickers": stock_tickers,
        "start_date": start_date,
        "end_date": end_date,
        "days_in_future": days_in_future
    }
    response = requests.post(url, json=data)
    if response.status_code == 200:
        return response.json()['predictions']
    else:
        raise Exception(f"Error fetching predictions: {response.text}")

def add_to_wallet(user_id, token, tickers, investment_amount, expected_return):
    url = f'http://localhost:8080/api/wallet/{user_id}'
    headers = {
        'Authorization': f'Bearer {token}'
    }
    tickers_data = [{
        "ticker": ticker,
        "amount_invested": investment_amount * top_5_weights_percent[ticker] / 100,
        "predictions": []
    } for ticker, weight in tickers.items()]
    data = {
        "tickers": tickers_data,
        "expected_gain": {}
    }
    print(f"Sending data to wallet: {data}")
    response = requests.post(url, json=data, headers=headers)
    if response.status_code == 200:
        print("Successfully added to wallet")
    else:
        print(f"Error adding to wallet: {response.text}")

if __name__ == "__main__":
    if len(sys.argv) != 7:
        print("Usage: python test.py <username> <password> <user_id> <investment_amount> <risk_profile> <investment_objectives>")
        sys.exit(1)

    username = sys.argv[1]
    password = sys.argv[2]
    user_id = sys.argv[3]
    investment_amount = float(sys.argv[4])
    risk_profile = sys.argv[5]
    investment_objectives = sys.argv[6]

    print(f"Creating portfolio for user ID: {user_id} with investment amount: {investment_amount}, risk profile: {risk_profile}, and investment objectives: {investment_objectives}")
    print(f"Login with Username: {username} and Password: {password}")

    token = login(username, password)
    
    try:
        # Parametri per la richiesta di previsione
        stock_tickers = ["AAPL", "GOOG", "MSFT", "AMZN", "TSLA", "META", "NFLX", "NVDA", "INTC", "ORCL","AMD", "IBM", "CSCO", "PYPL", "ADBE", "CRM", "ABNB", "UBER", "SQ", "SNAP","ZM", "PINS", "FSLY"]
        
        start_date = "2020-01-01"
        end_date = (datetime.today() - timedelta(days=1)).strftime('%Y-%m-%d')
        
        # Definizione dello switch-case per days_in_future in base a investment_objectives
        def get_days_in_future(investment_objectives):
            switcher = {
                'Short Term 1-5 Years': 730,
                'Medium Term 5-10 Years': 2556,
                'Long Term 10+ Years': 4017
                # Aggiungi altri casi secondo necessità
            }
            return switcher.get(investment_objectives, 30)  # Default a 30 giorni se non corrisponde a nessun caso

        days_in_future = get_days_in_future(investment_objectives)
        # Ottenere le previsioni dal servizio Flask
        predictions = get_predictions(stock_tickers, start_date, end_date, days_in_future)

        # Estrarre i rendimenti attesi dalle previsioni, ignorando i ticker con errori
        expected_returns_input = {}
        for ticker, preds in predictions.items():
            if isinstance(preds, list) and len(preds) > 0 and 'yhat' in preds[0]:
                initial_price = preds[0]['yhat']
                final_price = preds[-1]['yhat']
                # Calcolo del rendimento atteso come variazione percentuale media
                expected_returns_input[ticker] = (final_price - initial_price) / initial_price
            else:
                print(f"Skipping ticker {ticker} due to error or no valid predictions.")

        # Stampare i rendimenti attesi originali
        print("Rendimenti attesi originali:", expected_returns_input)

        # Scaricare dati storici per calcolare la matrice di covarianza
        def get_historical_data(stock_tickers, start_date, end_date):
            data = yf.download(stock_tickers, start=start_date, end=end_date)['Close']
            return data

        # Calcolo dei rendimenti attesi e della matrice di covarianza
        historical_data = get_historical_data(stock_tickers, start_date, end_date)
        expected_returns = pd.Series(expected_returns_input)
        cov_matrix = historical_data.pct_change().cov()

        # Funzione per calcolare il rendimento e il rischio del portafoglio
        def portfolio_performance(weights, expected_returns, cov_matrix):
            returns = np.dot(weights, expected_returns)
            risk = np.sqrt(np.dot(weights.T, np.dot(cov_matrix, weights)))
            return returns, risk

        # Funzione obiettivo per la minimizzazione del rischio
        def minimize_risk(weights, expected_returns, cov_matrix, target_return):
            returns, risk = portfolio_performance(weights, expected_returns, cov_matrix)
            return risk

        # Vincoli e limiti per l'ottimizzazione
        num_assets = len(expected_returns)

        # Definire il rendimento target desiderato
        def get_target_return(risk_profile):
            switcher = {
                'Spericolato': 0.14,
                'Moderato': 0.12,
                'Prudente': 0.10,
                'Molto Prudente': 0.8
                # Aggiungi altri casi secondo necessità
            }
            return switcher.get(risk_profile, 0.1)  # Default a 30 giorni se non corrisponde a nessun caso

        target_return = get_target_return(risk_profile)*(days_in_future/365)


        args = (expected_returns, cov_matrix, target_return)
        constraints = (
            {'type': 'eq', 'fun': lambda weights: np.sum(weights) - 1},
            {'type': 'eq', 'fun': lambda weights: portfolio_performance(weights, expected_returns, cov_matrix)[0] - target_return},
            {'type': 'ineq', 'fun': lambda weights: weights - 0.01},  # Limite inferiore per ogni peso
            {'type': 'ineq', 'fun': lambda weights: 0.3 - weights}    # Limite superiore per ogni peso
        )
        bounds = tuple((0, 1) for asset in range(num_assets))

        # Ottimizzazione del portafoglio
        result = minimize(minimize_risk, num_assets*[1./num_assets,], args=args,
                        method='SLSQP', bounds=bounds, constraints=constraints)

        # Risultati dell'ottimizzazione
        optimal_weights = result.x
        optimal_return, optimal_risk = portfolio_performance(optimal_weights, expected_returns, cov_matrix)

        # Identificazione delle migliori 5 pesature
        optimal_weights_series = pd.Series(optimal_weights, index=expected_returns.index)

        if(investment_amount<1500001):
            top_5_weights = optimal_weights_series.nlargest(5)
        else:
            top_5_weights = optimal_weights_series.nlargest(10) #distribuire maggiormente la liquidità con cifre alte,
                                                                #si potrebbe fare a scaglioni fino a 20 stock

        # Calcolo delle percentuali delle migliori 5 pesature
        top_5_weights_percent = (top_5_weights / top_5_weights.sum()) * 100

        # Funzione per formattare i numeri in modo leggibile
        def format_decimal(value, decimals=4):
            return f"{value:.{decimals}f}"

        # Stampa dei risultati con formattazione decimale
        print("Pesature ottimali (tutte):")
        print(optimal_weights_series.apply(lambda x: format_decimal(x)))

        print("Rendimento atteso del portafoglio:", format_decimal(optimal_return))
        print("Rischio del portafoglio:", format_decimal(optimal_risk))

        print("\nMigliori 5 pesature:")
        print(top_5_weights.apply(lambda x: format_decimal(x)))

        print("\nPercentuali delle migliori 5 pesature:")
        print(top_5_weights_percent.apply(lambda x: format_decimal(x)))
        
        


        # Aggiungi al wallet
        add_to_wallet(user_id, token, top_5_weights, investment_amount, format_decimal(optimal_return))

        # Visualizzazione della frontiera efficiente
        def efficient_frontier(expected_returns, cov_matrix, num_portfolios=100):
            results = np.zeros((3, num_portfolios))
            weights_record = []

            for i in range(num_portfolios):
                weights = np.random.random(len(expected_returns))
                weights /= np.sum(weights)
                returns, risk = portfolio_performance(weights, expected_returns, cov_matrix)
                results[0,i] = risk
                results[1,i] = returns
                results[2,i] = returns / risk
                weights_record.append(weights)

            return results, weights_record

        results, weights_record = efficient_frontier(expected_returns, cov_matrix)
        plt.figure(figsize=(10, 6))
        plt.scatter(results[0,:], results[1,:], c=results[2,:], cmap='viridis')
        plt.colorbar(label='Sharpe Ratio')
        plt.xlabel('Rischio (Deviazione Standard)')
        plt.ylabel('Rendimento Atteso')
        plt.scatter(optimal_risk, optimal_return, marker='*', color='r', s=200)  # Portafoglio ottimale
        plt.title('Frontiera Efficiente')
        plt.show()
    except Exception as e:
        print(f"An error occurred: {e}")
        sys.exit(1)
