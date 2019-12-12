import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ImportlistComponent } from './importlist.component';

describe('ImportlistComponent', () => {
  let component: ImportlistComponent;
  let fixture: ComponentFixture<ImportlistComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ImportlistComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ImportlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
