terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
  required_version = ">= 1.0"
}

provider "aws" {
  region = "eu-west-1"
}

# Use default VPC and subnet
data "aws_vpc" "default" {
  default = true
}

data "aws_subnet" "default" {
  filter {
    name   = "vpc-id"
    values = ["vpc-019168a95f36cba3f"]
  }

  filter {
    name   = "tag:Name"
    values = ["default"]
  }
}

# Security Group allowing inbound HTTP (port 80)
resource "aws_security_group" "web_sg" {
  name        = "blog-sg"
  description = "Allow HTTP & HTTPS inbound traffic"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTPS"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# EC2 instance
resource "aws_instance" "blog" {
  ami                         = data.aws_ami.ubuntu.id
  instance_type               = "t3.micro"
  subnet_id                   = data.aws_subnet.default.id
  vpc_security_group_ids      = [aws_security_group.web_sg.id]
  associate_public_ip_address = true

  user_data = <<-EOF
              #!/bin/bash
              apt-get update -y
              apt-get install -y docker.io curl

              # Install Docker Compose v2 (as plugin)
              mkdir -p /usr/local/lib/docker/cli-plugins
              curl -SL https://github.com/docker/compose/releases/download/v2.27.0/docker-compose-linux-x86_64 \
                   -o /usr/local/lib/docker/cli-plugins/docker-compose
              chmod +x /usr/local/lib/docker/cli-plugins/docker-compose

              systemctl start docker
              systemctl enable docker

              # Download and run the compose file
              mkdir -p /opt/blog
              curl -o /opt/blog/compose.yaml https://raw.githubusercontent.com/cyanrad/blog/refs/heads/master/compose.yaml
              cd /opt/blog
              docker compose up -d
              EOF

  tags = {
    Name = "go-webserver-instance"
  }
}

# Elastic IP for static IP
data "aws_eip" "blog_eip" {
  filter {
    name   = "tag:Name"
    values = ["blog-ip"]
  }
}

# Get latest Ubuntu AMI
data "aws_ami" "ubuntu" {
  most_recent = true
  owners      = ["099720109477"] # Canonical

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }
}

output "instance_public_ip" {
  value       = aws_eip.blog_eip.public_ip
  description = "Static Elastic IP of the EC2 instance"
}
