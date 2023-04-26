from locust import HttpUser, task

class SxMonitorDemoAppHighLatency(HttpUser):
    @task(1)
    def request_5s(self):
        self.client.get(f"/site/delay?sleep=5")
