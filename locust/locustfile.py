from locust import HttpUser, task

class SxMonitorDemoApp(HttpUser):
    @task(3)
    def request_200(self):
        self.client.get("/site/200")

    @task
    def request_503(self):
        self.client.get("/site/503")

    @task
    def request_404(self):
        self.client.get("/site/404")

    @task
    def request_40s(self):
        self.client.get(f"/site/delay?sleep=10")