<p align="center">
  <img  width="128" height="128" src="https://user-images.githubusercontent.com/37074372/154437448-3e82e0fc-ec70-4a32-b640-2184eeac99c4.png" alt="logo"/>
</p>



>**well_aware** is a cross platform osint tool written on go + [Fyne GUI](https://github.com/fyne-io/fyne). \
>Currently the tool contains three sections: *Domain Recon*, *Keyword Search* and *Custom Dorks*.
> <img width="758" alt="Screenshot 2022-01-26 at 21 06 43" src="https://user-images.githubusercontent.com/37074372/154436533-1e73e807-f018-424b-b311-940bf8a7a4a6.PNG"> \
>The **Domain Recon** section aimed to help in quick getting of information about target domain. \
>\
><img width="756" alt="Screenshot 2022-01-26 at 21 07 15" src="https://user-images.githubusercontent.com/37074372/154436656-82ccb8e3-464b-4e48-bb19-e918ccbc2085.PNG">\
>The **Keyword Search** utilizes google dorks and other sources to find some juicy information. \
>\
><img width="756" alt="Screenshot 2022-01-26 at 21 07 15" src="https://user-images.githubusercontent.com/37074372/154436795-f0482eea-dd89-47ab-85d5-6d830366fc5a.PNG">\
>The **Custom Dorks** section allows you to upload txt file with set of google dorks.

## Requirements

- Go 1.17+

## Installation

```sh
go get github.com/indigo-sadland/well_aware/
```
or you can build from binary
```sh
git clone https://github.com/indigo-sadland/well_aware/
cd well_aware
go build -o <OUTPUT_FILE_NAME> main.go
```

## Update

To update the tool use the following command:
```sh
go get -u github.com/indigo-sadland/well_aware/...
```

## Contribution
Please, feel free to send me requests on extending the tool :octocat:
