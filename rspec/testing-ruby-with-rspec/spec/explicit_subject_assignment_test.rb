RSpec.describe Array do 
  subject(:sally) do 
    ["ab", "cd"]
  end

  it 'correct array behavior' do 
    expect(sally.length).to eq(2)
    sally.pop 
    expect(sally.length).to eq(1)
  end

  it 'equal to original array' do 
    expect(sally).to eq(["ab", "cd"])
  end
end