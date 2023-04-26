from locust import HttpUser, task

class SxMonitorDemoAppHighBandwith(HttpUser):
    @task(3)
    def request_200(self):
        self.client.get("/site/high-bandwidth")