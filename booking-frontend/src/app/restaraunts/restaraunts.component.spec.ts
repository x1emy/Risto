import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RestarauntsComponent } from './restaraunts.component';

describe('RestarauntsComponent', () => {
  let component: RestarauntsComponent;
  let fixture: ComponentFixture<RestarauntsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RestarauntsComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RestarauntsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
