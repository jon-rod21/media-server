fetch('/api/library')
	.then(res => res.json())
	.then(files => {
		const container = document.getElementById('library');
		files.forEach(name => {
			const item = document.createElement('div');
			item.textContent = name;
			item.style.cursor = 'pointer';
			item.onclick = () => {
				const player = document.getElementById('player');
				player.src = '/stream/' + name;
				player.play();
			};
			container.appendChild(item);
		});
	});
