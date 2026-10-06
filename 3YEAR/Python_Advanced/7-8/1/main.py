import json 
import os
from datetime import datetime
from payments import CreditCard, PayPal 

FILE_NAME = "3YEAR/Python_Advanced/7-8/history.json"

def process_order(payment_method, amount):
    payment_method.pay(amount)

    record = {
        "method": type(payment_method).__name__,
        "amount": amount
    }
    if hasattr(payment_method, "balance"):
        record["balance"] = payment_method.balance
    record["time"] = datetime.now().strftime("%H:%M:%S")

    history = []
    if os.path.exists(FILE_NAME):
        with open(FILE_NAME, "r", encoding="utf-8") as file:
            try:
                history = json.load(file)
            except json.JSONDecodeError:
                history = []

    history.append(record)
    with open(FILE_NAME, "w", encoding="utf-8") as file:
        json.dump(history, file, ensure_ascii=False, indent=4)

card = CreditCard(100000)
paypal = PayPal()

process_order(card, 1000)
process_order(paypal, 1000)
process_order(card, 500)
process_order(paypal, 500)