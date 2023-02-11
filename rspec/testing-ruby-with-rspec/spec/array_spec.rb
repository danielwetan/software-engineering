require 'rspec'

RSpec.describe Array do
  it 'should start off empty' do 
    puts subject 
    puts subject.class 
    expect(subject.length).to eq(0)
  end

  it 'should contain an item' do 
    subject.push("Something")
    expect(subject.length).to eq(1)
  end
end