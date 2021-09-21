import json
import emoji
import qrcode

# our file (you are welcome!)
data = "https://youtu.be/XqZsoesa55w/"

# our file name
filename = "zoOapp.png"

# go to link
img = qrcode.make(data)

# save img to file
img.save(filename)

from flask import Flask, request

app = Flask(__name__)
	
# accept GET & POST methods
@app.route("/", methods=['GET', 'POST'])

def json_example():

	# receive JSON
	content = request.get_json(force=True
    
	# let emoji support
	animal = content["animal"]
	sound = content["sound"]
	count = str(content["count"])
	emoji =  {'Dragon':'\N{Dragon}', 'Sauropod':'\N{Sauropod}', 'T-Rex':'\N{T-Rex}'}
	emoji_icon = emoji.get(animal)

	# let's go fun
	if emoji_icon is None:
		return  "\nfilename\n"
	i=0
	res=''
	while i<int(count):
		temp = emoji_icon + animal + " says  "  +  sound + "\n"
		res =  temp + res
		i=i + 1
	res = "\n" + res + "\n Made with by \u1F996 Allozavrr \u1F996\n"
	return  res

if __name__ == '__main__':
app.run(host='0.0.0.0', port = 80, ssl_context = ('cert.pem', 'key.pem'))
