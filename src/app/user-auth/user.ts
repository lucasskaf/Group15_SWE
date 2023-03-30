export interface Movie {
    title: string,
    director: string,
    imglink: string,
    runtime: number,
    avgrating: number,
    providers: string[],
    databaseid: number
}


export interface User {
    "username" : string,
    "password" : string   
}

// rating: number,
// subscriptions: string[],
//email: string,
// genres: string[],
//watchlist: Movie[]