while (true) {
	const album = prompt("Album name: ");
	if (album?.toLowerCase() == "quit") window.close();

	const artist = prompt("Artist name: ");
	if (artist?.toLowerCase() == "quit") window.close();

	console.log(makeAlbum(String(artist), String(album)));
	console.log();
}


/**
 * @param {string} artist
 * @param {string} title
 * @returns {Object<string,string>}
 */
function makeAlbum(artist, title) {
	return { artist, title };
}
