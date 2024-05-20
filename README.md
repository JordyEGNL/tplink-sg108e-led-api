# tplink-sg108e-led-api
Simple API written in Python and Flask to turn the LEDs on the switch on and off

## WHY?
I have a TP-Link SG108E switch and I wanted to turn off the LEDs at night. The switch has a built in way to turn off the LEDs via the web interface, but I wanted to automate this process and turn off the LEDs at a specific time.
My Home Assistant will call this API to turn off the LEDs at a specific time. And turn them back on in the morning.

## Usage
```bash
docker compose up -d --build
```

## Endpoints
- `POST /led_off`
- `POST /led_on`

## Example
### Turn off the LEDs
```bash
curl -X POST http://localhost:5000/led_off -H "Content-Type: application/json" -d '{}'
```

### Turn on the LEDs
```bash
curl -X POST http://localhost:5000/led_on -H "Content-Type: application/json" -d '{}'
```

