export class UserModel{
    UserName?:string
    Role?:string
    Password?:string
    Permissions?:Permission[]
}

export class Permission{
    Name:string
    API:string
}


let User:UserModel={
    UserName:"",
    Role:"",
    Password:"",
    Permissions:[]
}

export default User