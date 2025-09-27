"""
LLD :


Real-Time Chat & Collaboration Tool

Design a backend system like Slack/MS Teams where:

1. Users can send direct messages or group messages
2. Messages can be text, image, or file
3. Presence (online/offline) is tracked in real-time
4. Notifications are sent on mentions


Approach: 
Entities: 
1. Users : Name, handle, status, active (online/offline), bio etc, phone_no, emails etc.. 
2. Message : meta -> type : text or images or files or video (url) 
3. User -> 1:1 Mapping DM 
        -> 1: M user_id --> group_id or channel_ids 
4. Groups/Channels: 
5. If message contains handle then trigger notifications. 
6. Notifications 
    - user_ids 
    - device_ids -> logged 
    - fcm_token --> mobile, laptop etc..  



Entities: 
Function/Class interfaces 
Schema 
Communication async/sync : 
Choose DB : 

MongoDB: ttl to dump if 2 years last message Expose write / reads 


TABLES: 
1. Users
    id 
    name 
    phone_number
    email_ids 
    status --> online 
    bio 
    created_at
    updated_at 
    is_active  


2. AuditLogs: 
    1. --> Meta events // perisents 
    2. 

3. UserGroupMessage: M:M 
    id  
    user_id
    group_id 
     

2. Groups:
    id 
    group_name 
    created_at 
    update_at 
    is_active  




APIs --> 

1. api/v1/user-> post create --> { all user, email} 
    api/v1/user/status -> post endpoint : online/offline M
    api/v1/user -> GET endpoints headers: Authorization --> User details 
    api/v1/user -> PATCH : name : patch 
    @vikaspal@kumar
    api/v1/send_message -> POST: content, tags = [handles, @vikaspal@kumar] -> successful 200 send. 
    api/v1/message -> patch 
      -               DElete 
    
    Groups: 
    api/v1/create-groups -> POST : { groupname} 
    api/v1/add-users/mebers -> Post : {group_ids, list_of_members } , PATCH, DELETE ENDPOINTS 


"""

def electronic_tax 
    return price 



calculate(1000, 10) $ 100
class User:

    def __init__(self, name, handle, email, phone_number, is_online=False): 
        self.name = name 
        self.handle = handle
        self.email = email
        self.phone_number = phone_number
        self.is_online = is_online
        self.bio = None

    def mark_online(self): 
          pass 
    
    def mark_offline(self):
          pass 
    
    def update_handle(self):
          pass 
    
    def update_bio(self):
          pass 
    
    def update_status(self): 
          pass 


class Database: 
     
     pass  


// Creation design pattern: factory pattern 
class MySQL(Database) 
    pass 

class MongoDB(Database):
     pass 
 
class Message: 
    """
        1. Two casess 
        2. send message 
    """
    def __init__(self, sender: User, receiver: User, content: str, msg_type: str):
          pass 
    
    def send_message(self): 
        pass 
    
 

class DirectMessage(Message): 
      def __init(self):
            pass 
      
      def send_message(self, sender_id, recever_id):
            print("Direct message")
       
class GroupMessage(Message): 

      def send_message(self, sender_id, group_id):
            print("trigger a group message")
            return super().send_message()
      


class Notification: 
    pass 


class Groups: 
    #  """
    #     - create a create 
    #     - add to a member or list of members 
    #     - remove a member from a groups 
    #     - set a name of group or channel 
    #  """
    def __init__(self):
        self.name = None 
        self.users = []
        self.owner = None 

    def add_member(self, users): 
        pass 
        # self.users.extend(users)

    def remove_member(self, users):
        pass

