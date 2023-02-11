RSpec.describe 'before and after hooks' do 
  # Reference to RSpec hooks
  # https://rspec.info/documentation/3.12/rspec-core/RSpec/Core/Hooks.html

  # before context hook
  before(:context) do 
    puts 'Before context'
  end

  after(:context) do 
    puts 'After context'
  end
  
  # before hook
  # running before example test block
  # for example seeding the database
  before(:example) do 
    puts 'Before example'
  end

  # after hook
  # running after example test block
  # for example reset the database
  after(:example) do 
    puts 'After example'
  end

  it 'example 1' do 
    expect(5 * 4).to eq(20)
  end

  it 'example 2' do 
    expect(3 - 2).to eq(1)
  end
end
