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

	async function runDrawInCanvas() {
		drawInCanvas();
	}

	async function runClickByMouse(e) {
		console.log(e);
		clickByMouse(e.offsetX, e.offsetY);
	}

	async function runKeyDown(e) {
		console.log(e);
		keyDown();
	}

	async function runKeyUp(e) {
		console.log(e);
		keyUp();
	}
	var resizeTimerId = -1;
	function resetCanvasSize() {
		canvas.Width = window.innerWidth;
		canvas.Height = window.innerHeight;
		if (resizeTimerId >= 0) clearTimeout(resizeTimerId);
		resizeTimerId = setTimeout(() => {
			//in next rendering event when resizing is finished.
			var canvas = document.querySelector('#canvas');
			setCanvas(window.innerWidth, window.innerHeight, {
				getContext: () => {
					return canvas.getContext('2d');
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
