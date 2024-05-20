from flask import Flask, request, jsonify
from flask_cors import CORS
import random
import requests

app = Flask(__name__)
CORS(app)

greetings = [
    "Hello! Welcome to our software company. How can I assist you today?",
    "Hi there! It's great to see you. Are you ready to explore our innovative software solutions?",
    "Greetings! I'm here to help you find the perfect products to meet your needs. How can I assist?",
    "Hello! How can I help you discover the best software products for your projects?",
    "Welcome! Let's dive into our software solutions. How can I be of service today?"
]

closings = [
    "Goodbye! If you have any more questions about our products, feel free to reach out.",
    "See you later! Remember, our team is always here to help you with your needs.",
    "Take care! Don't hesitate to contact us for further assistance.",
    "Farewell! Looking forward to helping you again.",
    "Bye! Keep innovating, and let us know if you need any more support."
]

@app.route('/chat', methods=['POST'])
def chat():
    data = request.json
    print(f"Message received: {data}")
    if not data or 'prompt' not in data:
        return jsonify({'message': 'No prompt provided' }), 400

    prompt = data['prompt']
    messages = data['messages']
    beginner_prompt =   """
                        As an AI assistant for our software company, your goal is to help users find the best software product from our offerings that meets their needs. Here are the products we offer:
                        1. ProManage - A project management tool with features like task tracking, Gantt charts, team collaboration, and time management. Suitable for small to large teams. Medium price.
                        2. ClientConnect - A CRM system designed to manage customer relationships, sales pipelines, and marketing campaigns. Integrates with popular email services and social media platforms. High price.
                        3. AccountEase - An accounting software that handles invoicing, payroll, expense tracking, and financial reporting. Available for both small and larger businesses. Medium price.
                        4. TeamSync - A communication and collaboration tool featuring chat, video conferencing, file sharing, and integration with other productivity apps. Ideal for remote teams. Small price.
                        5. MarketGuru - A marketing automation platform with tools for email marketing, social media management, analytics, and campaign management. Designed for marketing teams of all sizes. High price.
                        When assisting users, consider the following criteria to recommend the best product:
                        1. The type of software they are looking for (e.g., project management, CRM, accounting).
                        2. Specific features or functionalities that are important to them.
                        3. The scale of use (e.g., individual, small team, large organization).
                        4. Any integration needs with other tools or software they might have.
                        Based on this information, recommend the most suitable product from our offerings and provide a detailed explanation of why each product fits their needs.
                        If information is lacking, ask for more. Don't ask many questions at once. Always write short response at maximum 120 words. Always respond in “you” perspective as talking directly to the user.
                        Here is the user message:
                    """
    prompt = beginner_prompt + prompt

    url = 'http://localhost:11434/api/generate'
    payload = {
        'model': 'llama3',
        'prompt': prompt,
        'messages': messages,
        'stream': False
    }
    headers = {'Content-Type': 'application/json'}

    try:
        response = requests.post(url=url, json=payload, headers=headers)
        response.raise_for_status()
        result = response.json()
        message_response = result.get('response', 'No response received')
        return jsonify({'prompt': prompt,'message': message_response })

    except requests.exceptions.RequestException as e:
        print(f'Error: {e}')
        return jsonify({'message': 'Sorry, an error occurred while responding.' }), 500

@app.route('/greet', methods=['GET'])
def greet():
    greeting = random.choice(greetings)
    return jsonify({'message': greeting})

@app.route('/close', methods=['GET'])
def close():
    closing = random.choice(closings)
    return jsonify({'message': closing})

if __name__ == '__main__':
    app.run(debug=True, port=8080)
