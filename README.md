# MiniChatApplication-Golang-vue-js-socketio

# steps to start the application
# start the golang server using : go run main.go
# cd client : npm run serve


# each client should have a unqiue name
# check whether any room is created if not create on and assign the client
# one room should contain only two clients
# Separate room should be created for every two clients
# if any client disconnected , they will be removed from the room and
#  immediate next new client will be added to the room which has only one client
# both the clients should have a username before starting to chat
# send email will be enabled once chat starts
# if any new client joined to the room where the previous client left they can get the previous chat in email
# user will be able to update the profile pic
# send the profile pic along with msg
# show scroll for msg container alone
# show username left on groups on client left and leave group