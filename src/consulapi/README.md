## 打包

```bash
make all
```

## 封装不同的架构的Docker

```bash
FROM scratch

WORKDIR /
ADD consulapi_amd64 /consulapi # 根据需求,将不同的架构的包打进docker

CMD ["/consulapi"]
```

## k8s运行

```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: consulapi
  namespace: devops
spec:
  replicas: 1
  revisionHistoryLimit: 3
  selector:
    matchLabels:
      app: consulapi
  template:
    metadata:
      labels:
        app: consulapi
    spec:
      containers:
      - name: consulapi
        image: harbor.zjhw.com/library/consulapi
        imagePullPolicy: Always
        env:
        # consul server ip
        - name: CONSULHOST
          value: 192.168.66.100 
        resources:
          requests:
            memory: "200Mi"
            cpu: "200m"
          limits:
            memory: "400Mi"
            cpu: "400m"
        ports:
        - name: consulapi
          containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: consulapi
  namespace: devops
spec:
  selector:
    app: consulapi
  ports:
  - name: consulapi
    port: 8080
    targetPort: 8080

---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: consulapi
  namespace: devops
spec:
  rules:
  - host: consulapi.sysadmin.com
    http:
      paths:
      - path: /
        backend:
          serviceName: consulapi
          servicePort: 8080
```

## 访问地址

```bash
http://consulapi.sysadmin.com/swagger/index.html
```