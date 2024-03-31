# Basic CI-CD Test

setting up a basic CI/CD pipeline for a web application. The application consists of a simple static website and a basic RESTful API.

## Tasks

1. Version Control (Git):

    - [x] Create a Git repository for any test web application.
    - [x] Commit the initial codebase including the static website and any sample test API.

2. CI/CD Pipeline

    - [x] Choose a CI/CD tool (e.g., GitHub Actions, Jenkins).
      - _Github Actions_
    - [x] Set up a basic pipeline that triggers on each commit to the develop branch.
      - _Implemented for both develop and main branch_
    - [x] Include stages for building and deploying both the static website and API.
      - _Stages implemented:_
        - _Testing(for backend only)_
        - _Build and push image to Dockerhub_
        - _Deploy Containers to Azure_

3. Automated Testing:

    - [x] Implement a simple automated test for the API (e.g., a basic endpoint response check).
      - _Tests written in Go_
    - [x] Ensure that the CI/CD pipeline fails if the test fails.
      - _Tested to fail if the test cases fail_

4. Deployment:

    - [x] Deploy the static website to a simple web server.
      - _Firebase deployment done_
      - _Nginx server container deployment to azure_
    - [x] Deploy the API to a server, ensuring that it can handle basic HTTP requests.
      - _Containerized API Server deployed to azure_

### Project Structure

[Architecture diagram - Draw.io](https://drive.google.com/file/d/1qU1IhZAexsxA4jlTcJkS5mX_Tzvqf1ez/view?usp=sharing "Architecture diagram")

### Deployments

- Frontend
  - [Frontend Dev Branch Deployment](https://dev-basic-ci-frontend.wonderfulwave-531b1711.centralus.azurecontainerapps.io "frontend Dev")
  - [Frontend Main Branch Deployment](https://basic-ci-frontend.wonderfulwave-531b1711.centralus.azurecontainerapps.io "frontend Main")

  - [Firebase based Deployment](https://dev-basic-ci.firebaseapp.com/)

- Backend
  - [Backend Dev Branch Deployment](https://dev-basic-ci-cd-backend.wonderfulwave-531b1711.centralus.azurecontainerapps.io "backend Dev")
  - [Backend Main Branch Deployment](https://basic-ci-cd-backend.wonderfulwave-531b1711.centralus.azurecontainerapps.io "backend")
