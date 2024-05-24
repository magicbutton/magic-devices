FROM mcr.microsoft.com/powershell

# RUN pwsh -c "Install-Module -Name PnP.PowerShell -Force -AllowPrerelease -Scope AllUsers;" 

RUN apt update -y
RUN apt upgrade -y
RUN apt install golang-1.21 -y
RUN apt install unzip -y
ENV GOBIN="/usr/local/bin"
ENV PATH="/usr/lib/go-1.21/bin:${PATH}"

ENV KITCHEN_HOME="/kitchens"
RUN go install github.com/koksmat-com/koksmat@v2.1.4.18
RUN koksmat context init sharepoint
WORKDIR /kitchens
COPY ./.koksmat/kitchenroot .
WORKDIR /kitchens/magic-devices
COPY . .  

# Installs dependencies 
RUN koksmat ship init
RUN koksmat ship install

# Builds the app
WORKDIR /kitchens/magic-devices/.koksmat/app
RUN go install




CMD [ "sleep","infinity"]