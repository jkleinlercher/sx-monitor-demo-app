from locust import HttpUser, task

class SxMonitorDemoAppHighMemory(HttpUser):
    @task(3)
    def request_200(self):
        self.client.get("/site/high-memory")