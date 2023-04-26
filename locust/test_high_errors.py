from locust import HttpUser, task

class SxMonitorDemoAppHighErrors(HttpUser):
    @task(2)
    def request_200(self):
        self.client.get("/site/200")

    @task(10)
    def request_503(self):
        self.client.get("/site/503")

    @task(10)
    def request_404(self):
        self.client.get("/site/404")

