from locust import HttpUser, task

class SxMonitorDemoApp(HttpUser):
    @task(10)
    def request_200(self):
        self.client.get("/site/200")

    @task(2)
    def request_503(self):
        self.client.get("/site/503")

    @task(2)
    def request_404(self):
        self.client.get("/site/404")

    @task(3)
    def request_2s(self):
        self.client.get(f"/site/delay?sleep=2")

    @task(1)
    def request_5s(self):
        self.client.get(f"/site/delay?sleep=5")
