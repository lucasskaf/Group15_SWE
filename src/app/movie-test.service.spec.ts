import { TestBed } from '@angular/core/testing';

import { MovieTestService } from './movie-test.service';

describe('HelloWorldService', () => {
  let service: MovieTestService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(MovieTestService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
