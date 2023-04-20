import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MoviePopupComponent } from './movie-popup.component';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Emmiters } from '../emitters/emmiters';
import { FormBuilder } from '@angular/forms';
import { of } from 'rxjs';
import { Movie } from '../user-auth/user';
import { DialogModule } from 'primeng/dialog';

describe('MoviePopupComponent', () => {
  let component: MoviePopupComponent;
  let fixture: ComponentFixture<MoviePopupComponent>;
  let movieCompService: MovieGeneratorService
  let httpClientSpy: jasmine.SpyObj<HttpClient>
  let USER = {
    username: "own"
  }

  beforeEach(async () => {
    httpClientSpy = jasmine.createSpyObj('HttpClient', ['post', 'get'])
    movieCompService = new MovieGeneratorService(httpClientSpy)

    await TestBed.configureTestingModule({
      declarations: [ MoviePopupComponent ],
      imports: [HttpClientModule, DialogModule],
      providers: [MovieGeneratorService, FormBuilder]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MoviePopupComponent);
    component = fixture.componentInstance;
    movieCompService = TestBed.inject(MovieGeneratorService)
    fixture.detectChanges();
  })

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should emit isPopupOpen event with false value when closeMovie is called', () => {
    const moviePopupComponent = new MoviePopupComponent(new MovieGeneratorService(httpClientSpy), new FormBuilder());
    spyOn(Emmiters.isPopupOpen, 'emit');
    moviePopupComponent.closeMovie();
    expect(Emmiters.isPopupOpen.emit).toHaveBeenCalledWith(false);
  });

  it('should fetch movie posts successfully when generateMoviePosts is called', () => {
    const movieGeneratorService = new MovieGeneratorService(httpClientSpy);
    const moviePopupComponent = new MoviePopupComponent(movieGeneratorService, new FormBuilder());
    spyOn(movieGeneratorService, 'getMoviePosts').and.returnValue(of([{
      id: 1,
      movie_id: "123",
      username: "test",
      title: "Test Post",
      body: "This is a test post."
    }]));
    moviePopupComponent.generatedMovie.id = 123;
    moviePopupComponent.generateMoviePosts(moviePopupComponent.generatedMovie.id.toString());
    expect(moviePopupComponent.moviePosts.length).toBe(1);
    expect(moviePopupComponent.didFoundPosts).toBe(true);
  });

  it('should add a movie to watchList and emit isMovieWatched event with true value when addToWatchlist is called', () => {
    const movieGeneratorService = new MovieGeneratorService(httpClientSpy);
    const moviePopupComponent = new MoviePopupComponent(movieGeneratorService, new FormBuilder());
    spyOn(movieGeneratorService, 'addToWatchList').and.returnValue(of({
      id: 1,
      title: "Test Movie",
      vote_average: 7.5
    }));
    spyOn(Emmiters.isMovieWatched, 'emit');
    spyOn(Emmiters.watchList, 'push');
    const testMovie: Movie = {
      id: 1,
      title: "Test Movie",
      vote_average: 7.5
    };
    moviePopupComponent.addToWatchlist(testMovie);
    expect(movieGeneratorService.addToWatchList).toHaveBeenCalledWith(testMovie);
    expect(moviePopupComponent.isMovieWatched).toBe(true);
    expect(Emmiters.isMovieWatched.emit).toHaveBeenCalledWith(true);
    expect(Emmiters.watchList.push).toHaveBeenCalledWith({
      id: 1,
      title: "Test Movie",
      vote_average: 7.5
    });
  });
});
