# Inspiration from https://github.com/psmode/essstat/blob/master/essstat.py

from flask import Flask, request, jsonify
import requests
import os

app = Flask(__name__)

def login_and_control_led(ip, username, password, action):
    login_url = f"http://{ip}/logon.cgi"
    led_url = f"http://{ip}/led_on_set.cgi?rd_led={action}&led_cfg=Apply"

    # Start a session to maintain cookies
    session = requests.Session()

    # Login payload
    login_data = {
        "logon": "Login",
        "username": username,
        "password": password
    }

    # Headers for login
    login_headers = {
        'Referer': f'http://{ip}/Logout.htm',
        'Content-Type': 'application/x-www-form-urlencoded'
    }

    try:
        # Login to the switch
        login_response = session.post(login_url, data=login_data, headers=login_headers)
        login_response.raise_for_status()

        # Headers for controlling the LED
        led_headers = {
            'Referer': f'http://{ip}/'
        }

        # Control the LED
        led_response = session.get(led_url, headers=led_headers)
        led_response.raise_for_status()

        return {"status": "success", "action": "on" if action == "1" else "off"}

    except requests.exceptions.RequestException as e:
        return {"status": "error", "message": str(e)}

@app.route('/led_on', methods=['POST'])
def led_on():
    data = request.json
    ip = data.get('ip') or os.getenv('TP_LINK_IP')
    username = data.get('username') or os.getenv('TP_LINK_USERNAME')
    password = data.get('password') or os.getenv('TP_LINK_PASSWORD')
    
    if not ip or not username or not password:
        return jsonify({"status": "error", "message": "Missing required parameters"}), 400

    result = login_and_control_led(ip, username, password, "1")
    return jsonify(result)

@app.route('/led_off', methods=['POST'])
def led_off():
    data = request.json
    ip = data.get('ip') or os.getenv('TP_LINK_IP')
    username = data.get('username') or os.getenv('TP_LINK_USERNAME')
    password = data.get('password') or os.getenv('TP_LINK_PASSWORD')
    
    if not ip or not username or not password:
        return jsonify({"status": "error", "message": "Missing required parameters"}), 400

    result = login_and_control_led(ip, username, password, "0")
    return jsonify(result)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
