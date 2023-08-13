all_messages = ["Hello there", "Readability counts", "Less is more", "Simple is better"]
sent_messages = list(str())


def send_messages(messages: list[str], sent: list[str]):
    while len(messages) > 0:
        sent.append(messages[0])
        all_messages.pop(0)


def show_messages(messages: list[str]):
    for msg in messages:
        print(msg)


send_messages(all_messages, sent_messages)
print(all_messages)
print(sent_messages)
