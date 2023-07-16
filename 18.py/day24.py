import re
import copy

a = """Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with
 an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning,
 slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack
 that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire,
 cold) with an attack that does 12 slashing damage at initiative 4"""
a = open("day24.txt").read()

class Team:
    def __init__(self, desc, name, boost=0):
        groupstr = desc.split("\n")
        self.groups = {i+1: Group(g, name+" group "+str(i+1)) for i, g in enumerate(groupstr)}
        self.name = name
        for g in self.groups.values():
            g.dmg += boost

    def print_status(self):
        print(self.name+":")
        for group_id, group in self.groups.items():
            print("Group",group_id,"contains",group.units,"units")


    def get_units(self):
        return sum([g.units for g in self.groups.values()])

    def print_final(self):
        print(self.name, "remaining units", self.get_units())

    def has_units(self):
        return self.get_units() > 0

    def target_selection(self, other):
        targets = copy.deepcopy(other.groups)
        ordered_groups = sorted(self.groups.items(), key = lambda g: g[1].getPowerIniSort(), reverse=True)
        for group_id, group in ordered_groups:
            if len(targets) == 0: break #happens if all targets popped
            target_id, target = max(targets.items(), key=lambda t: t[1].get_damage_taken_eff_pow_ini(group.get_effective_power(), group.dmgtype)) #get best target
            log_dmg = target.get_damage_taken(group.get_effective_power(), group.dmgtype) #damage you'd do, just for logging
            if log_dmg == 0: continue # if you do 0 damage, don't pick that target
            #print(self.name,"group",group_id,"would deal defending group", target_id, log_dmg, "damage")
            group.next_target = other.groups[target_id] # set next target
            targets.pop(target_id) # remove from options

    def cleanup(self):
        # filter dict like https://stackoverflow.com/questions/2844516/how-to-filter-a-dictionary-according-to-an-arbitrary-condition-function
        # kick out no-unit-groups
        self.groups = {k: v for k, v in self.groups.items() if v.units > 0}

class Group:
    def __init__(self, desc, name):
        self.name = name
        #print(desc)
        numbers = re.findall(r"\d+", desc)
        numbers = [int(nr) for nr in numbers]
        self.units, self.hp, self.dmg, self.ini = numbers
        self.weak = re.findall(r"weak to (.+)\)", desc)
        if self.weak:
            # hotfix for parsing issue
            self.weak = re.sub(r";.+","", self.weak[0])
            self.weak = self.weak.split(", ")
        self.immune = re.findall(r"immune to (.+)\)", desc)
        if self.immune:
            self.immune = re.sub(r";.+","", self.immune[0])
            self.immune = self.immune.split(", ")
        self.dmgtype = re.findall(r"does \d+ (.+) damage", desc)[0]
        self.next_target = None
        #print(self)

    def get_effective_power(self):
        return self.units * self.dmg

    def __repr__(self):
        return str(["name", self.name, \
                        "units",self.units, "hp",self.hp, \
                        "dmg",self.dmg, "ini",self.ini, \
                        "weak",self.weak, "immune",self.immune, "dmgtype",self.dmgtype])

    # needed for sorting who attacks first
    def getPowerIniSort(self):
        return (self.get_effective_power(), self.ini)

    # needed for sorting who gets targeted first
    def get_damage_taken_eff_pow_ini(self, effectivePower, dmgtype):
        dt = self.get_damage_taken(effectivePower, dmgtype)
        return (dt, self.get_effective_power(), self.ini)

    def get_damage_taken(self, effectivePower, dmgtype):
        # By default, an attacking group would deal damage equal to its effective power to the defending group. 
        # However, if the defending group is immune to the attacking group's attack type, the defending group 
        # instead takes no damage; if the defending group is weak to the attacking group's attack type, 
        # the defending group instead takes double damage.
        if dmgtype in self.weak:
            return effectivePower*2
        elif dmgtype in self.immune:
            return 0
        else:
            return effectivePower

    def attack(self):
        if self.next_target == None:
            return
        dmg = self.next_target.get_damage_taken(self.get_effective_power(), self.dmgtype)
        # int division. 531 dmg only kill 5 units with 100 health
        killed_units = dmg // self.next_target.hp
        killed_units = min(killed_units, self.next_target.units)
        self.next_target.units -= killed_units
        #print(self.name,"attacks",self.next_target.name,"killing",killed_units,"units")
        self.next_target, self.next_damage = None, 0 # so we don't attack a team without selecting it

immu, infec = a.strip().replace("Immune System:\n","").replace("\n ", " ").split("\n\nInfection:\n")
#print(immu,"\n",infec)
boost = 69
print(boost)
immuT = Team(immu, "Immune system", boost)
infecT = Team(infec, "Infection")
last_units = 0

while immuT.has_units() and infecT.has_units():
    # status phase    
    #immuT.print_status()
    #infecT.print_status()
    print()

    # target phase
    infecT.target_selection(immuT)
    immuT.target_selection(infecT)
    #print()

    # attack phase
    # get team with biggest ini first
    tgroups = [*infecT.groups.values(), *immuT.groups.values()]
    tgroups = sorted(tgroups, key=lambda g: g.ini, reverse=True)
    for group in tgroups:
        group.attack()
    #print()

    # cleanup phase: delete groups with 0 units
    infecT.cleanup()
    immuT.cleanup()

    units = infecT.get_units() + immuT.get_units()
    if last_units == units:
        print("no progress made, exiting this try")
        break
    last_units = units

infecT.print_final()
immuT.print_final()
# toolow 14847