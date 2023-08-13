def send_messages(messages: list[str], sent: list[str]):
    while len(messages) > 0:
        sent.append(messages[0])
        messages.pop(0)


def show_messages(messages: list[str]):
    for msg in messages:
        print(msg)


messages = ["Hello there", "Readability counts", "Less is more", "Simple is better"]
sent_messages = list(str())

send_messages(messages, sent_messages)
print(messages)
print(sent_messages)
