class ProgrammingLanguage
  attr_reader :name 
  
  def initialize(name = "Ruby")
    @name = name
  end
end


RSpec.describe ProgrammingLanguage do 
  let(:language) { ProgrammingLanguage.new('Python') }

  it 'should store the name of the language' do 
    expect(language.name).to eq('Python')
  end

  # overwritting let spec
  context 'with no argument' do 
    # available only in nested context
    let(:language) { ProgrammingLanguage.new }

    it 'should default to Ruby as the name' do 
      expect(language.name).to eq('Ruby')
    end
  end
end
