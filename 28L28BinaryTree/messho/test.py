"""
Inventory Management System
Develop a program to manage the inventory for an ecommerce company. In this ecommerce website once the user initiates his payment, inventory will be blocked for him for 5min. If the user finishes the payment and comes back within 5min the blocked inventory will be allocated to him else the inventory will be released back into the app for everyone else to order.To support the above features implement the following methods. Store all the data in memory, in appropriate data structures.

/*create product with given productId, name and inventory count*/  => void addProduct(String productId, String name, Integer count)

/*return the available quantity for given product*/  =>  Integer getInventory(String productId)

/*Will be called when the supplier wants to update inventory for a product.Update the inventory of given product with the given count*/  =>  boolean updateInventory(String productId, Integer count)

/*Will be called when the user initiates payment for an order. Block the inventory for the given product and for the given order reference for 5min*/  =>  boolean blockInventory(String productId, Integer count, String orderId)

/*Will be called when the user completes payment for his order.Reduce the ordered quantity permanently for the product corresponding to given orderId. If this method is not called within 5min from blockInventory, inventory should be released back*/  =>  boolean confirmOrder(String orderId) ttl redis -> pubsub  
"""
# block
# updateInventory 
""" 
inventory_product_map . 
concurrently locked on shared 
product_id quantity  blocked_quanity 
maggie      5         0 
  

app-> s2s -> checkout->maggie -> product_id 121 

M1-> m2 -> in-memory 

r1, r2 , r3 
2   2 2 
Atomic  
lock()
blocked_quanity; 
 2 
quanitiy 
 7 
unblock()


placing order for 2 maggie add_to_carts 

block_inventory_orders -> 

order_id product_id  quantity  expired_times 
   12         121    2      time.now() + 5 mins  -> updateInventory() 

   1. functional code  for demo 
   2. 

""" 
from typing import Dict
import threading 
import time

class Inventory:  

    def __init__(self): 
        self.inventory : Dict [str, int]  = {} 
        self.lock = threading.Lock()
        self.blockedInventory : Dict[str, Dict[str, int]] = {}

 

    def add_product(self, product_name, count):    
        with self.lock: 
            if product_name in self.inventory: 
                self.inventory[product_name] += count   
                print(f" updated inventory") 
            else:
                self.inventory[product_name] = count 
                print("add product")


    def get_inventory(self): 
        with self.lock:
            return dict(self.inventory)
    
    def get_inventory_product_name(self, product_name):  
        if product_name in self.inventory:
            return self.inventory.get(product_name,0) 
        return False


    def update_inventory(self, product_id, count): 
        with self.lock:
            if product_id in  self.product_quantity: 
                self.product_quantity[product_id] += count 
                return True
            return False

    def block_inventory(self, order_id, product_name, quantity, block_time=5): 
        with self.lock:
            available = self.inventory.get(product_name, 0) 
            if available < quantity: 
                return False 
            self.inventory[product_name] -= quantity 

            order_data = self.blockedInventory.get(order_id) 
            if not order_data:  
                self.blockedInventory[order_id] = {} 
                order_data = self.blockedInventory[order_id ]
            prev_qty = self.blockedInventory[order_id].get(product_name, 0)
            self.blockedInventory[order_id][product_name] = prev_qty + quantity 
            timer = threading.Timer(block_time, self.release_block, args=(order_id,))
            timer.daemon = True 
            timer.start()
            return True  
        
    def release_block(self, order_id): 
        with self.lock:
            blocked_items = self.blockedInventory.pop(order_id, None) 
            if not blocked_items:  
                print("order is confirmed")
                return  
            
            for product_name , q in blocked_items.items():  
                self.inventory[product_name] = self.inventory.get(product_name, 0) + q 


    def confirm_order(self, order_id):
        with self.lock:
            if order_id not in self.blockedInventory: 
                return  False 
            blocked = self.blockedInventory.pop(order_id) 
            return True  
       




if __name__ == "__main__": 
    inm = Inventory() 
    inm.add_product("maggie", 10)     
    print("Inventory by productName", inm.get_inventory_product_name("maggie"))
    print("Inventory by productName", inm.get_inventory_product_name("Iphone"))
    print("Current Inventory:", inm.get_inventory())     
    inm.block_inventory("OR1", "maggie", 3)     
    print("Current Inventory:", inm.get_inventory())    
    print("order confirmation", inm.confirm_order("OR1"))   
    print("Current Inventory:", inm.get_inventory()) 



