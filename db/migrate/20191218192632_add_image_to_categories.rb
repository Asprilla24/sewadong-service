class AddImageToCategories < ActiveRecord::Migration[5.2]
  def change
    add_column :categories, :image, :string, limit: 255
  end
end
