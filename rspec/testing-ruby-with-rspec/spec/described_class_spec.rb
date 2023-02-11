class King 
  attr_reader :name 

  def initialize(name)
    @name = name
  end
end

RSpec.describe King do 
  subject { described_class.new('Boris') }
  # is equal to 
  # subject { King.new('Boris') }
  let(:louis) { described_class.new('Louis') }

  it 'represent a great person' do 
    expect(subject.name).to eq('Boris')
    expect(louis.name).to eq('Louis')
  end
end

# Reference to rspec described_class
# https://relishapp.com/rspec/rspec-core/docs/metadata/described-class

