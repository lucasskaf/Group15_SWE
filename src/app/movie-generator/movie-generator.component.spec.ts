import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormBuilder, FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MovieGeneratorComponent } from './movie-generator.component';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { Emmiters } from '../emitters/emmiters';
import { of } from 'rxjs';
import { MoviePopupComponent } from '../movie-popup/movie-popup.component';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatSliderModule} from '@angular/material/slider';
import { DialogModule } from 'primeng/dialog';
import { By } from '@angular/platform-browser';
import { Movie } from '../user-auth/user';

describe('MovieGeneratorComponent', () => {
  let component: MovieGeneratorComponent;
  let fixture: ComponentFixture<MovieGeneratorComponent>;
  let movieGeneratorServiceSpy: jasmine.SpyObj<MovieGeneratorService>;

  beforeEach(async () => {
    const spy = jasmine.createSpyObj('MovieGeneratorService', ['getRandomMovieWithFilters']);

    await TestBed.configureTestingModule({
      declarations: [ MovieGeneratorComponent, MoviePopupComponent ],
      providers: [
        { provide: MovieGeneratorService, useValue: spy }
      ],
      imports: [
        FormsModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatSelectModule,
        MatSliderModule,
        DialogModule
      ]
    })
    .compileComponents();

    movieGeneratorServiceSpy = TestBed.inject(MovieGeneratorService) as jasmine.SpyObj<MovieGeneratorService>;
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(MovieGeneratorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should call getGeneratedMovie with correct argument', () => {
    const filters = {
      actors: ['Tom Hanks', 'Brad Pitt'],
      genres: 'comedy',
      min_rating: 6,
      max_runtime: 120,
      streaming_providers: ['Netflix', 'Hulu'],
    };
    const generatorData = {
      actors: 'Tom Hanks, Brad Pitt',
      genres: 'comedy',
      min_rating: '3',
      max_runtime: '120',
      streaming_providers: 'Netflix, Hulu',
    };

    spyOn(component, 'getGeneratedMovie');
    component.onSubmit(generatorData);

    expect(component.getGeneratedMovie).toHaveBeenCalledWith(filters);
  });

  it('should call getGeneratedMovie function when submitting form', () => {
    spyOn(component, 'getGeneratedMovie').and.callThrough();
    const formValues = {
      actors: 'Tom Hanks, Matt Damon',
      genres: 'Action, Drama',
      min_rating: '4',
      max_runtime: '120',
      streaming_providers: 'Netflix, Amazon Prime',
    };
    component.onSubmit(formValues);
    expect(component.getGeneratedMovie).toHaveBeenCalledWith(formValues);
  });
})