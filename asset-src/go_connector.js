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
	let mod, inst;
	window.InitFinished = false;
	WebAssembly.instantiateStreaming(fetch('result.wasm'), go.importObject)
		.then(async result => {
			mod = result.module;
			inst = result.instance;
			inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance

			// document.getElementById('runButton').disabled = false;
			console.clear();
			window.InitFinished = true;
			await go.run(inst);
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
			console.log('set - canvas ');
			setCanvas(window.innerWidth, window.innerHeight, {
				getContext: () => {
					return canvas.getContext('2d');
				},
				getBackground: fileName => {
					const image = new Image();
					console.log('drawing back:', fileName);
					image.src = `./images/${fileName}.jpeg`;

					const ctx = canvas.getContext('2d');
					image.addEventListener('load', () => {
						console.log(image);
						setBackground({ image });
						ctx.drawImage(
							image,
							0,
							0,
							image.naturalWidth,
							image.naturalHeight,
							0,
							0,
							window.innerWidth,
							window.innerHeight
						);
					});
					// image.onload = img => {

					// 	ctx.drawImage(image, 0, 0, 100, 100);
					// };

					return image;
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
