from firebase_admin import auth

user = auth.get_user(uid)
print 'Successfully fetched user data: {0}'.format(user.uid)