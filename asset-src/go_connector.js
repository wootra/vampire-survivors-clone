'use strict';

(() => {
	if (!WebAssembly.instantiateStreaming) {
		// polyfill
		WebAssembly.instantiateStreaming = async (resp, importObject) => {
			const source = await (await resp).arrayBuffer();
			return await WebAssembly.instantiate(source, importObject);
		};
	}

	const go = new Go();
	const imageSet = {};
	const images = {
		'back-1': 'back-1.jpeg',
		fish: 'fish.png',
		cat: 'cats.png',
	};
	const loadedImages = {
		loaded: 0,
		total: Object.keys(images).length,
	};

	function loadImages() {
		Object.keys(images).forEach(key => {
			const image = new Image();

			image.src = `./images/${images[key]}`; //later, it will be changed to imagesets

			image.addEventListener('load', () => {
				imageSet[key] = image;
				loadedImages.loaded++;
			});
		});
	}
	loadImages();

	let mod, inst;
	window.InitFinished = false;
	WebAssembly.instantiateStreaming(fetch('result.wasm'), go.importObject)
		.then(async result => {
			mod = result.module;
			inst = result.instance;
			inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance

			// document.getElementById('runButton').disabled = false;
			// console.clear();
			window.InitFinished = true;
			await go.run(inst);
			setGlueFunctions({
				getLoadStatus: () => {
					return loadedImages;
				},
			});
		})
		.catch(err => {
			console.error(err);
		});

	var resizeTimerId = -1;

	function resetCanvasSize() {
		var canvas = document.querySelector('#canvas');
		canvas.width = window.innerWidth;
		canvas.height = window.innerHeight;

		console.log('reset canvas size');
		if (resizeTimerId >= 0) clearTimeout(resizeTimerId);
		resizeTimerId = setTimeout(() => {
			//in next rendering event when resizing is finished.
			canvas = document.querySelector('#canvas');
			const ctx = canvas.getContext('2d');
			const halfW = window.innerWidth / 2;
			const halfH = window.innerHeight / 2;
			document.createElement('canvas', {
				width: window.innerWidth,
				height: window.innerHeight,
			});

			ctx.translate(halfW, halfH);
			console.log('set - canvas ');
			setCanvas(window.innerWidth, window.innerHeight, {
				getContext: () => {
					const ctx = canvas.getContext('2d');
					return ctx;
				},

				getBackground: fileName => {
					const img = imageSet[fileName];
					setBackground({ image: img });
					ctx.drawImage(
						img,
						0,
						0,
						img.naturalWidth,
						img.naturalHeight,
						-halfW,
						-halfH,
						window.innerWidth,
						window.innerHeight
					);
					return img;
				},
				getCharacterImage: (fileName, frame, x, y, size) => {
					const img = imageSet[fileName];
					const xd = img.naturalHeight * frame;
					ctx.drawImage(
						img,
						xd,
						0,
						img.naturalHeight,
						img.naturalHeight,
						x,
						y,
						size,
						size
					);
					return img;
				},
			});
			resizeTimerId = -1;
		}, 100);
	}

	addEventListener('resize', e => {
		// console.log(e);
		resetCanvasSize();
	});

	window.addEventListener('DOMContentLoaded', event => {
		console.log('dom is loaded...');
		var intervalId = setInterval(() => {
			if (window.InitFinished) {
				resetCanvasSize();
				clearInterval(intervalId);
				setGlueFunctions({
					getLoadStatus: () => {
						return loadedImages;
					},
				});
			} else {
				console.log('interval...');
			}
		}, 1000);
	});

	window.addEventListener('keydown', async event => {
		// console.log(event);
		keyDown(event.code);
	});
	window.addEventListener('keyup', async event => {
		// console.log(event.code);
		keyUp(event);
	});
})();
