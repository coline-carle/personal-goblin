INSERT INTO items (id, watch_auctions)
  VALUES
    -- primordial
    (152668, FALSE), -- Expulsom

    -- herbalism --
    (152505, TRUE), --  riverbud
    (152506, TRUE), -- star moss
    (152507, TRUE), -- akunda's bite
    (152508, TRUE), -- winter's kiss
    (152509, TRUE), -- sirren pollens
    (152510, TRUE), -- anchor weed
    (152511, TRUE),  -- sea-stalk

    -- alchemy --
    -- equipement
    (152632, TRUE), -- Surging Alchemist Stone
    (152637, TRUE), -- Siren's Alchemist Stone
    (152634, TRUE), -- Endless Tincture of Renewed Combat
    (152636, TRUE), -- Endless Tincture of Fractional Power

    -- heal / mana potions
    (152494, TRUE), -- Coastal Healing Potion
    (152495, TRUE), -- Coastal Mana Potion
    (163082, TRUE), -- Coastal Rejuvenation Potion

    -- flasks
    (152638, TRUE), -- Flask of the Currents
    (152639, TRUE), -- Flask of Endless Fathoms
    (152640, TRUE), -- Flask of the Vast Horizon
    (152641, TRUE), -- Flask of the Undertow

    -- cauldron
    (162519, FALSE), -- Mystical Cauldron

    -- batle potions
    (163222, TRUE), -- Battle Potion of Intellect
    (163223, TRUE), -- Battle Potion of Agility
    (163224, TRUE), -- Battle Potion of Strength
    (163225, TRUE), -- Battle Potion of Stamina
    (152559, TRUE), -- Potion of Rising Death
    (152560, TRUE), -- Potion of Bursting Blood
    (152561, TRUE), -- Potion of Replenishment
    (152557, TRUE), -- Steelskin Potion

    -- utility potions
    (162113, FALSE), -- Potion of Herb Tracking
    (152496, TRUE), -- Demitri's Draught of Deception
    (152497, TRUE), -- Lightfoot Potion
    (152503, TRUE), -- Potion of Concealment
    (152550, TRUE), -- Sea Mist Potion

    -- transmutation
    (152578, FALSE), -- Sack of Herbs
    (160325, FALSE), -- Quivering Sac
    (152580, FALSE), -- Pile of Cloth
    (152581, FALSE), -- Bag of Jewels
    (152582, FALSE), -- Stack of Skins
    (160322, FALSE) -- Pile of Ore
  ON CONFLICT (id)
  DO UPDATE
    SET watch_auctions =  EXCLUDED.watch_auctions;
