import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ArticleTileComponent } from './article-tile.component';

describe('ArticleTileComponent', () => {
  let component: ArticleTileComponent;
  let fixture: ComponentFixture<ArticleTileComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ArticleTileComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ArticleTileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
