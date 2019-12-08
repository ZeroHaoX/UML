// 出货
export class ExportRe{
    eid?:string
    eprice?:number
    ecount?:number
    etotalprice?:number
    buyer?:string
    detial?:string
    imdate?:string
    gname?:string
    shipper?:string
    bphone?:string
    profit?:number
    edate?:number
}

// 进货
export class ImportRe{
    id?:string
    gname?:string
    imprice?:number
    imcount?:number
    imtotalprice?:number
    shipper?:string
    sphone?:string
    detial?:string
    imdate?:string
}