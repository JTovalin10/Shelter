import subprocess
import sys
import os
import signal
import time

ROOT = os.path.dirname(os.path.abspath(__file__))
BACKEND = os.path.join(ROOT, "backend")
FRONTEND = os.path.join(ROOT, "frontend")

processes = []

def shutdown(sig, frame):
    print("\nShutting down...")
    for p in processes:
        p.terminate()
    sys.exit(0)

signal.signal(signal.SIGINT, shutdown)

print("Starting backend...")
backend = subprocess.Popen(
    ["go", "run", "./cmd/server/main.go"],
    cwd=BACKEND
)
processes.append(backend)

print("Starting frontend...")
frontend = subprocess.Popen(
    ["npm", "run", "dev"],
    cwd=FRONTEND
)
processes.append(frontend)

print("\n frontend: http://localhost:5173")
print("  backend: http://localhost:8080")
print("\nCtrl+C to stop both\n")

for p in processes:
    p.wait()
