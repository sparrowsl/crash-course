console.log(makeAlbum("Don Moen", "Thank you Lord", 1));
console.log(makeAlbum("Adele", "Go easy on me", 30));
console.log(makeAlbum("Jon Bellion", "I feel it", 1));


/**
 * @param {string} artist
 * @param {string} title
 * @param {number} songs
 * @returns {Object<string, string|number>}
 */
function makeAlbum(artist, title, songs) {
	return { artist, title, songs };
}
