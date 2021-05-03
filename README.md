This project is based on features for app+chat room in each app.

Design Schema:
    1. App that works as a container for the project. Multiple apps can be deployed on the same server, with each one having its owd
       configuration
    2. Each app will have its main extension file that'll act as the entry points for the app.
    3. Each app will have its own room list.
    4. Each room in an app will have an event handler and can have its own userlist.
    5. User will have to login to app to be able to continue further and join rooms

Models:
Events:
    - Login
    - Logout
    - Join App
    - Subscribe App
    - Join Room
    - Leave Room
    - Disconnect
    - Reconnect
App
    - App Handler
    - Event Handlers
    - Room List
    - Users

User
    - id
    - username
    - roomlist

Room
    - Event Handlers
    - Users