name := "e-shop-api"
version := "1.0-SNAPSHOT"

lazy val root = (project in file(".")).enablePlugins(PlayScala)

scalaVersion := "3.3.0"

libraryDependencies ++= Seq(
  guice
)