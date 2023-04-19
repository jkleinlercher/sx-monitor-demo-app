from locust import HttpUser, task

class SxMonitorDemoAppHighCPU(HttpUser):
    @task(3)
    def request_200(self):
        self.client.get("/site/high-cpu")