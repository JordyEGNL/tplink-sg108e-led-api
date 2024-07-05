# tplink-sg108e-led-api
Simple API written in Go to turn the LEDs on my switch on and off

## WHY?
I have a TP-Link SG108E switch and I wanted to turn off the LEDs at night. The switch has a built in way to turn off the LEDs via the web interface, but I wanted to automate this process and turn off the LEDs at a specific time.
My Home Assistant will call this API to turn off the LEDs at a specific time. And turn them back on in the morning.

## Usage
### docker-compose.yml file
```yaml
services:
  tplink-api:
    container_name: tplink-api
    image: ghcr.io/jordyegnl/tplink-sg108e-led-api:latest
    ports:
      - "5000:5000"
    environment:
      TP_LINK_IP: 192.168.1.1
      TP_LINK_USERNAME: admin
      TP_LINK_PASSWORD: admin
```

### Run the container
```bash
docker compose up -d
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

