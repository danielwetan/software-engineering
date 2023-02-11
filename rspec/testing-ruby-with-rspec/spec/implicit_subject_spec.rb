RSpec.describe Hash do 
  # RSpec documentation about implicit subject
  # https://relishapp.com/rspec/rspec-core/v/3-8/docs/subject/implicitly-defined-subject

  # Difference between RSpec subject and let 
  # https://stackoverflow.com/questions/38437162/whats-the-difference-between-rspecs-subject-and-let-when-should-they-be-used

  # implicit subject in rspec
  # let(:subject) { Hash.new }

  it 'should start off empty' do 
    puts subject 
    puts subject.class  
    subject[:some_key] = "Some value"
    expect(subject.length).to eq(1)
  end

  it 'is isolated between example' do 
    expect(subject.length).to eq(0)
  end
end