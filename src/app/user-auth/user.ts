export interface Movie {
    adult?: boolean , 
	backdrop_path?: string,
	budget?: number,
	genres?: string[],
	homepage?: string,
	id?: number,
	original_language?: string,
	original_title?: string,   
	overview?: string,   
	popularity?: number, 
	poster_path?: string, 
	production_companies?: string[],
	production_countries?: string[],
	release_date?: string, 
	revenue?: number,     
	runtime?: number,    
	spoken_languages?: string[],
	status?: string,  
	tagline?: string,  
	title?: string, 
	vote_average?: number,
	vote_count?: number,  
	user_rating?: number
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