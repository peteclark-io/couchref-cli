default: clean dev

clean:
	- rm ../public/js/*bundle.js

dev: pre-steps dependencies webpack serve

pre-steps: /usr/local/bin/webpack
/usr/local/bin/webpack:
	npm install -g webpack

dependencies: ./node_modules
./node_modules: ./package.json
	npm install

webpack:
	webpack --watch --progress --colors &

serve:
	node server.js
	open http://localhost:9001/

prod:
	webpack -p --config webpack.config.prod.js --progress --colors
