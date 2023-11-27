#See https://aka.ms/customizecontainer to learn how to customize your debug container and how Visual Studio uses this Dockerfile to build your images for faster debugging.


#DEFAULT CONFIG BY VISUAL STUDIO
#FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
#WORKDIR /app
#EXPOSE 80
#EXPOSE 443
#
#FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
#ARG BUILD_CONFIGURATION=Release
#WORKDIR /src
#COPY ["./distributedsystemsproject.csproj", "distributedsystemsproject/"]
#RUN dotnet restore "./distributedsystemsproject/./distributedsystemsproject.csproj"
#COPY . .
#WORKDIR "/src/distributedsystemsproject"
#RUN dotnet build "./distributedsystemsproject.csproj" -c $BUILD_CONFIGURATION -o /app/build
#
#FROM build AS publish
#ARG BUILD_CONFIGURATION=Release
#RUN dotnet publish "./distributedsystemsproject.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false
#
#FROM base AS final
#WORKDIR /app
#COPY --from=publish /app/publish .
##ENTRYPOINT ["dotnet", "distributedsystemsproject.dll"]


# CHAT-GPT

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /app
COPY *.csproj ./
RUN dotnet restore
COPY . ./
RUN dotnet publish -c Release -o out
FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS runtime
WORKDIR /app
COPY --from=build /app/out ./
EXPOSE 80
EXPOSE 443
CMD ["dotnet", "distributedsystemsproject.dll"]
ENTRYPOINT ["dotnet", "distributedsystemsproject.dll"]