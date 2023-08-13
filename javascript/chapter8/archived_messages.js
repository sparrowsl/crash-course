let allMessages = ["Hello there", "Readability counts", "Less is more", "Simple is better"];
let sentMessages = [];

sendMessages(allMessages, sentMessages);
console.log(allMessages);
console.log(sentMessages);


/**
 * @param {string[]} messages
 * @param {string[]} sent
 */
function sendMessages(messages, sent) {
	while (messages.length > 0) {
		sent.push(messages[0]);
		allMessages.shift();
	}
}

/** @param {string[]} messages */
function showMessages(messages) {
	for (const v in messages) console.log(v);
}
