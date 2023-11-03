from linepy import *
# line = LINE('611781024@crru.ac.th', '064689Arm')
line = LINE('22b60d0df15ae1025d189556b6502a38832c1029feccfdd7ab59490440af3dc6')
line.log("Auth Token : " + str(line.authToken))
line.log("Timeline Token : " + str(line.tl.channelAccessToken))