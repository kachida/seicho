# Seicho 🌱

**Seicho** means *growth* or *development* in Japanese.  
It’s a small terminal app that helps you stay organized, track your goals, and build daily consistency.  
Think of it as a personal‑growth companion that gives you gentle reminders and measures your progress over time.

---

## What Seicho Can Do

- Add, edit, or remove goals  
- Log your daily progress  
- See your success and consistency stats  
- Tag goals by category (like health, learning, creativity)  
- Get reminders if you forget to log today  
- Set simple alarms to take action  
- Read short inspirational quotes  
- Export your progress to a CSV file  
- Use “Zen mode” for a quick motivational boost

---

## Installation

1. Clone the project and build it:
   ```bash
   git clone https://github.com/<your-username>/seicho.git
   cd seicho
   go mod tidy
   go build -o seicho


---

## Getting started
Run the app for the first time:


    ```bash
    seicho
    ```

---

## How to use

### Add a goal

Create a new goal with a description, frequency, and tags:

    ```bash
    seicho add "learn elixir language" -d "30 min deliberate practice on elixir language" -f daily -t coding,functional programming
    ```

### List all goals

See everything you’re tracking:
    ```bash
    seicho list
    ```

### Log progress

Record what you did today:
    ```bash
    seicho log 1 done
    ```
or mark it missed
    ```bash
    seicho log 1 missed
    ```

### View stats

Check your success rate and consistency:
    ```bash
    seicho stats
    ```

### Edit or remove goals

Update or delete goals you no longer need:

    ```bash
    seicho edit 1 -t "Evening run" -d "Jog after work"
    seicho remove 2
    ```

---


