"""
    LLD Topics: 
    1. Encapsulation: Programs of hiding the internal state and requiring all interaction to be performed through an object's methods.

    Isolation Levels: how concurrent transactions are handles / interact with the databases 
    1. Read Uncommitted Example: Transaction A updates a row but hasn’t committed. Transaction B reads that update — if A rolls back, B saw invalid data.
        Dirty reads are there. 
    2. Read Committed 
        A Transactions reads only the commited data we have. Prevents Dirty reads kos : 
            Non-repeadables and phathom reads still exists. 
    3. Repeatable Read
    4. Serializable : prevents from phathom reads . 
"""

class BankAccount:

    def __init__(self, account_number, balance):
        self.__account_number = account_number  # Private attribute
        self.__balance = balance

    def deposit(self, amount):
        self.__balance += amount

    def withdraw(self, amount):
        if amount <= self.__balance:
            self.__balance -= amount
        else:
            print("Insufficient funds")

    def get_balance(self):
        return self.__balance