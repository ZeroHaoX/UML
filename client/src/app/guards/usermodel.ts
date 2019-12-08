export class UserModel{
    userName?:string
    role?:string
    // Password?:string
    permissions?:Permission[]
}

export class Permission{
    name?:string
    api?:string
}


let User:UserModel={
    userName:"",
    role:"",
    // Password:"",
    permissions:[]
}

export default User